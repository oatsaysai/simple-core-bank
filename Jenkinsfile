pipeline {
    agent any

    stages {
        stage('Pull From Git') {
            steps {
                git branch: 'main', url: 'https://github.com/oatsaysai/simple-core-bank.git'
            }
        }
        stage('Docker Build') {
            steps {
                sh 'docker build --rm -t docker.skrss.com:5000/gitops-webapp:${BUILD_NUMBER} .'
            }
        }
        stage('Push Docker Imahe') {
            steps {
                sh 'docker push docker.skrss.com:5000/gitops-webapp:${BUILD_NUMBER}'
            }
        }
        stage('Clear Previous Image') {
            steps {
                script{
                    try {
                        sh 'docker rmi -f $(docker images -q --filter "before=docker.skrss.com:5000/gitops-webapp:${BUILD_NUMBER}" docker.skrss.com:5000/gitops-webapp)'
                    } catch (err) {
                        echo "don't have any Previous Image"
                    }
                }
            }
        }
        stage('Edit ArgoCD') {
             steps {
                script {
                    withCredentials([gitUsernamePassword(credentialsId: 'build.user', gitToolName: 'Default')]) {
                        sh '''
                            rm -rf gitops-webapp
                            git clone https://repo.blockfint.com/sakkarin/gitops-webapp.git
                            cd gitops-webapp
                            sed -i "13s|newTag: .*$|newTag: \\"${BUILD_NUMBER}\\"|" deployment/dev/kustomization.yaml
                            git config --global user.email "sakkarin@blockfint.com"
                            git config --global user.name "Jenkins"
                            git commit -am "Updates kustomization.yaml with $BUILD_NUMBER"
                            git push -u origin master
                        '''
                    }
                }
            }
        }
        stage('Success') {
            steps {
                echo 'Deploy Successfully'
            }
        }
    }
}
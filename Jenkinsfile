pipeline{
    agent any

    environment{
        AWS_ACCOUNT_ID = 'aws-id'
        AWS_DEFAULT_REGION = 'us-east-1'
        BACKEND_ECR_REPOSITORY_NAME = 'backend-ecr-repo-name'
        FRONTEND_ECR_REPOSITORY_NAME = 'frontend-ecr-repo-name'       
        BACKEND_REPOSITORY_URI = 'backend-ecr-repo-uri'
        FRONTEND_REPOSITORY_URI = 'frontend-ecr-repo-uri'
        SCANNER_HOME = 'sonar-scanner'
    }

    stages{
        stage('Git checkout'){
            steps{
                sh 'echo Checkout passed'
            }
        }
        stage('Trivy Scan'){
            steps{
                sh 'trivy fs --format table -o fs-report.html .'
            }
        }

        stage('SonarQube'){
            steps{
                sh " $SCANNER_HOME/bin/sonar-scanner Dsonar.projectKey=URL-shortener Dsonar.ProjectName=URL-shortener"
            }
        }
        stage('Build & tag Docker image for backend'){
            steps{
                script{
                    withDockerRegistry(creadentialsId: 'docker-cred', toolName:'docker')
                    sh 'docker system prune -f'
                    sh 'docker container prune -f'
                    sh 'docker build -t ${REPOSITORY_URI}${BACKEND_ECR_REPOSITORY_NAME}::$(BUILD_NUMBER) .'

                }
            }
        }
        stage('Trivy Scan Backend Docker image'){
            steps{
                sh 'trivy image --format table -o fs-report.html $(BACKEND_ECR_REPOSITORY_NAME)'
            }
        }
        stage('Push to BACKEND_ECR'){
            steps{
                script{
                    withDockerRegistry(creadentialsId: 'docker-cred', toolName:'docker')
                    sh 'aws ecr get-login-password --region ${AWS_DEFAULT_REGION} | docker login --username AWS --password-stdin ${BACKEND_REPOSITORY_URI}'
                    sh 'docker push ${BACKEND_REPOSITORY_URI}${BACKEND_ECR_REPOSITORY_NAME}:${BUILD_NUMBER}'
                }
            }
        }

        stage('Trivy Scan Frontend Docker image'){
            steps{
                sh 'trivy image --format table -o fs-report.html $(FRONTEND_ECR_REPOSITORY_NAME)'
            }
        }
        stage('Push to FRONTEND_ECR'){
            steps{
                script{
                    withDockerRegistry(creadentialsId: 'docker-cred', toolName:'docker')
                    sh 'aws ecr get-login-password --region ${AWS_DEFAULT_REGION} | docker login --username AWS --password-stdin ${FRONTEND_REPOSITORY_URI}'
                    sh 'docker push ${FRONTEND_REPOSITORY_URI}${FRONTEND_ECR_REPOSITORY_NAME}:${BUILD_NUMBER}'
                }
            }
        }
        stage('Update Deployment File'){
            environment{
                GIT_REPO_NAME = 'URL-Shortner-With-GO'
                GIT_USER_REPO = 'ARUP-G'
            }
            steps{
                dir('kubernetes-Manifests-file/forntend'){
                    withCredentials([string(credentialsId: 'github', variable: 'GITHUB_TOKEN')]){
                        sh ```
                            git config user.email "darup2019@gmail.com"
                            git config user.name "ARUP-G"
                            BUILD_NUMBER=${BUILD_NUMBER}
                            echo $BUILD_NUMBER
                            imageTag=$(grep -oP '(?<=frontend:)[^ ]+' deployment.yml)
                            echo $imageTag
                            sed -i "s/${AWS_ECR_REPO_NAME}:${imageTag}/${AWS_ECR_REPO_NAME}:${BUILD_NUMBER}/" deployment.yml
                            git add deployment.yml
                            git commit -m "Update deployment Image to version \${BUILD_NUMBER}"
                            git push https://${GITHUB_TOKEN}@github.com/${GIT_USER_NAME}/${GIT_REPO_NAME} HEAD:master
                        ```
                    }
                }
            }
        }
        stage('Deploy eo EKS'){
            steps{
                script{
                    sh 'aws eks update-kubeconfig --region ${AWS_DEFAULT_REGION} --name uel-shortener-cluster'
                    sh 'kubectl apply -f kubernetes-Manifests-file/secret.yml'
                    sh 'kubectl apply -f kubernetes-Manifests-file/deployment.yml'
                    sh 'kubectl apply -f kubernetes-Manifests-file/service.yml'
                    sh 'kubectl apply -f kubernetes-Manifests-file/ingress.yml'
                }
            }
        }
        post{
            always{
                cleanWS() 
            }
        }
    }
}
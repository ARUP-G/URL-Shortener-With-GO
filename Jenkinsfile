pipeline{
    agent any

    environment{

        // AWS
        AWS_ACCOUNT_ID        = 'aws-id'
        AWS_DEFAULT_REGION    = 'us-east-1'
        SCANNER_HOME          = 'sonar-scanner'
        EKS_CLUSTER_NAME      = 'your-eks-cluster'

        // ECR
        BACKEND_ECR_REPOSITORY_NAME     = 'backend-ecr-repo-name'
        FRONTEND_ECR_REPOSITORY_NAME    = 'frontend-ecr-repo-name'       
        BACKEND_REPOSITORY_URI          = 'backend-ecr-repo-uri'
        FRONTEND_REPOSITORY_URI         = 'frontend-ecr-repo-uri'

        //git 
        GIT_USER_ACCOUNT    = 'git_user_account'
        GIT_USER_NAME       = 'git_user_name'
        GIT_REPO_NAME       = 'git_repo_name'

        // Helm Configuration
        HELM_RELEASE_NAME   = 'helm_relese_name'  
        HELM_CHART_PATH     = 'helm_chart_path'

        // Kubernetes Configuration
        K8S_NAMESPACE       = 'k8s_namespace'
    }

    stages{
        stage('Git checkout'){
            steps{
                script{
                    checkout([
                    $class: 'GitSCM',
                    branches: [[name: '*/main']],
                    userRemoteConfigs: [[url: 'https://github.com/ARUP-G/URL-Shortener-With-GO.git']]
                    ])
                    sh 'echo Checkout passed'
                } 
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
                    sh './backend/Dockerfile'
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
        stage('Build & tag Docker image for forntend'){
            steps{
                script{
                    withDockerRegistry(creadentialsId: 'docker-cred', toolName:'docker')
                    sh './frontend/Dockerfile'
                    sh 'docker build -t ${REPOSITORY_URI}${FRONTEND_ECR_REPOSITORY_NAME}::$(BUILD_NUMBER) .'
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
        stage('Update Deployment File') {
            steps {
                dir('kubernetes-Manifests-file/frontend') {
                    withCredentials([string(credentialsId: 'github', variable: 'GITHUB_TOKEN')]) {
                        script {
                            sh """
                            git config --global user.email "${GIT_USER_ACCOUNT}"
                            git config --global user.name "${GIT_USER_NAME}"
                            echo ${BUILD_NUMBER}
                            // Update backend image tag in values.yaml
                            sed -i 's|backendImage:.*|backendImage:|g' ${HELM_CHART_PATH}
                            sed -i 's|repository: .*|repository: ${BACKEND_REPOSITORY_URI}/${BACKEND_ECR_REPOSITORY_NAME}|g' ${HELM_CHART_PATH}
                            sed -i 's|tag: .*|tag: "${BUILD_NUMBER}"|g' ${HELM_CHART_PATH}
                            // Update frontend image tag in values.yaml
                            sed -i 's|frontendImage:.*|frontendImage:|g' ${HELM_CHART_PATH}
                            sed -i 's|repository: .*|repository: ${FRONTEND_REPOSITORY_URI}/${FRONTEND_ECR_REPOSITORY_NAME}|g' ${HELM_CHART_PATH}
                            sed -i 's|tag: .*|tag: "${BUILD_NUMBER}"|g' ${HELM_CHART_PATH}
                            git add ${HELM_CHART_PATH}/values.yaml
                            git commit -m "Update deployment Image to version ${BUILD_NUMBER}"
                            git push https://${GITHUB_TOKEN}@github.com/${GIT_USER_NAME}/${GIT_REPO_NAME} HEAD:main
                            """
                        }
                    }
                }
            }
        }
        stage('Deploy eo EKS'){
            steps{
                script{
                    sh 'aws eks update-kubeconfig --region ${AWS_DEFAULT_REGION} --name uel-shortener-cluster'
                    helm upgrade --install ${HELM_RELEASE_NAME} ${HELM_CHART_PATH} --namespace ${K8S_NAMESPACE} --create-namespace
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
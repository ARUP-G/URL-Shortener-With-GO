pipeline{
    agent any

    environment {

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
        
        // SonarQube Configuration
        SONARQUBE_URL       = 'http://your-sonarqube-url'
        SONARQUBE_TOKEN     = 'sonarqube-token'

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

        stage('SonarQube Scan') {
            steps {
                withCredentials([string(credentialsId: 'sonarqube-token', variable: 'SONARQUBE_TOKEN')]) {
                    sh """
                    $SCANNER_HOME/bin/sonar-scanner \
                        -Dsonar.projectKey=URL-shortener \
                        -Dsonar.projectName=URL-shortener \
                        -Dsonar.host.url=${SONARQUBE_URL} \
                        -Dsonar.login=${SONARQUBE_TOKEN}
                    """
                }
            }
        }

        stage('Build & tag Docker image for backend'){
            steps{
                script{
                    sh 'docker system prune -f'
                    sh 'docker container prune -f'
                    sh 'docker build -t ${REPOSITORY_URI}${BACKEND_ECR_REPOSITORY_NAME}::$(BUILD_NUMBER) -f ./backend/Dockerfile .'

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
                    // Login to AWS ECR
                    sh 'aws ecr get-login-password --region ${AWS_DEFAULT_REGION} | docker login --username AWS --password-stdin ${BACKEND_REPOSITORY_URI}'

                    // Push Docker image
                    sh 'docker push ${BACKEND_REPOSITORY_URI}:${BUILD_NUMBER}'
                }
            }
        }
        stage('Build & tag Docker image for forntend'){
            steps{
                script{
                    sh 'docker build -t ${REPOSITORY_URI}${FRONTEND_ECR_REPOSITORY_NAME}::$(BUILD_NUMBER) -f ./frontend/Dockerfile'
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
                        // Configure Git user
                        sh 'git config --global user.email "${GIT_USER_ACCOUNT}"'
                        sh 'git config --global user.name "${GIT_USER_NAME}"'

                        // Update backend image repository and tag in Helm values.yaml
                        sh """
                        sed -i 's|repository:.*public.ecr.aws/y6q4k1r8/prod-backend|repository: ${BACKEND_REPOSITORY_URI}|g' ${HELM_CHART_PATH}/values.yaml
                        sed -i 's|tag:.*|tag: "${BUILD_NUMBER}"|g' ${HELM_CHART_PATH}/values.yaml
                        """

                        // Update frontend image repository and tag in Helm values.yaml
                        sh """
                        sed -i 's|repository:.*public.ecr.aws/y6q4k1r8/prod-frontend|repository: ${FRONTEND_REPOSITORY_URI}|g' ${HELM_CHART_PATH}/values.yaml
                        sed -i 's|tag:.*|tag: "${BUILD_NUMBER}"|g' ${HELM_CHART_PATH}/values.yaml
                        """

                        // Optionally update mongo image tag if needed
                        //  Uncomment and modify if you want to update mongo image details as well
                        //  sh """
                        //  sed -i 's|repository: mongo|repository: ${MONGO_REPOSITORY_URI}|g' ${HELM_CHART_PATH}/values.yaml
                        //  sed -i 's|tag:.*|tag: "${BUILD_NUMBER}"|g' ${HELM_CHART_PATH}/values.yaml
                        //  """

                        // Check if there are changes to commit
                        def changes = sh(script: 'git status --porcelain', returnStdout: true).trim()
                        if (changes) {
                            sh 'git add ${HELM_CHART_PATH}/values.yaml'
                            sh 'git commit -m "Update deployment images to version ${BUILD_NUMBER}"'
                            sh 'git push https://${GITHUB_TOKEN}@github.com/${GIT_USER_NAME}/${GIT_REPO_NAME} HEAD:main'
                        } else {
                            echo 'No changes to commit.'
                        }
                        }
                    }
                }
            }
        }

        stage('Deploy on EKS'){
            steps{
                script{
                    sh 'aws eks update-kubeconfig --region ${AWS_DEFAULT_REGION} --name uel-shortener-cluster'
                    sh 'helm upgrade --install ${HELM_RELEASE_NAME} ${HELM_CHART_PATH} --namespace ${K8S_NAMESPACE} --create-namespace'
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

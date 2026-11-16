
pipeline {
    agent none

    stages {
        stage('Construir Docker en AWS') {
            agent {
                label 'built-in'
            }
            steps {
                checkout scm
                echo 'RED CLOUD - Construyendo Docker en AWS'
                sh 'docker build -t red-cloud-web:1.0 .'
            }
        }

        stage('Desplegar Docker en AWS') {
            agent {
                label 'built-in'
            }
            steps {
                echo 'Desplegando aplicacion en AWS'
                sh '''
                    docker rm -f red-cloud-web || true
                    docker run -d --name red-cloud-web -p 8081:8080 red-cloud-web:1.0
                '''
            }
        }

        stage('Construir imagen para Kubernetes') {
            agent {
                label 'kubernetes'
            }
            steps {
                checkout scm
                echo 'Construyendo imagen en Minikube'
                sh 'minikube image build -t red-cloud-web:${BUILD_NUMBER} .'
            }
        }

        stage('Desplegar en Kubernetes') {
            agent {
                label 'kubernetes'
            }
            steps {
                checkout scm
                echo 'Actualizando Kubernetes'
                sh '''
                    kubectl apply -f service.yaml
                    kubectl apply -f deployment.yaml
                    kubectl set image deployment/red-cloud-web red-cloud-web=red-cloud-web:${BUILD_NUMBER}
                    kubectl rollout status deployment/red-cloud-web --timeout=180s
                    kubectl get pods
                '''
            }
        }
    }
}

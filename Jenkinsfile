pipeline {
    agent { label 'built-in' }

    stages {
        stage('Construir imagen de pruebas') {
            steps {
                echo 'RED CLOUD - Rama pruebas'
                sh 'docker build -t red-cloud-web:pruebas-${BUILD_NUMBER} .'
            }
        }

        stage('Desplegar en pruebas (puerto 8082)') {
            steps {
                sh '''
                    docker rm -f red-cloud-web-pruebas || true
                    docker run -d --name red-cloud-web-pruebas -e ENTORNO=pruebas -p 8082:8080 red-cloud-web:pruebas-${BUILD_NUMBER}
                '''
            }
        }

        stage('Comprobar que responde') {
            steps {
                sh 'sleep 3 && curl -s http://localhost:8082/ | grep "Entorno"'
            }
        }
    }
}
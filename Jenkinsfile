
pipeline {
    agent any

    stages {
        stage('Comprobar GitHub') {
            steps {
                echo 'RED CLOUD - Inicio del Pipeline'
                sh 'ls -la'
            }
        }

        stage('Construir imagen Docker') {
            steps {
                echo 'Construyendo la aplicacion Go'
                sh 'docker build -t red-cloud-web:1.0 .'
            }
        }

        stage('Desplegar aplicacion') {
            steps {
                echo 'Desplegando la aplicacion en Docker'
                sh '''
                    docker rm -f red-cloud-web || true
                    docker run -d --name red-cloud-web -p 8081:8080 red-cloud-web:1.0
                '''
            }
        }
    }
}

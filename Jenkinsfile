
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
    }
}

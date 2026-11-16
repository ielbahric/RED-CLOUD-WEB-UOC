
pipeline {
    agent any

    stages {
        stage('Comprobar GitHub') {
            steps {
                echo 'RED CLOUD - Inicio del Pipeline'
                echo 'Codigo descargado desde GitHub correctamente'
            }
        }

        stage('Comprobar archivos') {
            steps {
                sh 'ls -la'
            }
        }
    }
}

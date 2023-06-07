# SERVICIO APP MOVIL UPLA

<!-- ![IMAGES DE GO LANG](images/ladder.svg) -->
<img src="src/helper/images/logo_upla.svg" alt="Imagen go" width="200" />

<font size="5">
Aplicación encargada de realizar de crear los recursos que van hacer usados desde la app movil.
</font>

## Iniciar

Esta proyecto esta echo en golang.
<img src="src/helper/images/ladder.svg" alt="Imagen go" width="50" />

Algunos recursos para iniciar con este proyecto puedes ver en:

- [Go](https://go.dev/) Go en un lenguaje de programación concurrente y compilado con tipado estático.
- [Docker](https://www.docker.com/) Docker es la tecnología en contenedores que permite crear y usar imagenes para desplegar aplicación.
- [Visual Studio](https://code.visualstudio.com/) Editor de código para todos tipos de lenguaje de programación.
- [Git](https://git-scm.com/) Software de control de versiones.
- [Git Hub](https://github.com/) Plataforma de alojamiento de proyecto de todo ámbito.

## Instalación

Siga los pasos para iniciar el desarrollo:

1. Clona el proyecto o agrague el ssh al repositorio para contribuir en nuevos cambios [Git Hub - UPLA SERVICIO APP MOVIL](https://github.com/luissince/servicio-app-movil-upla)

    1.1. Agregue por ssh para la integración

    #Code

        /** 
        ** Para el proceso de integración **
        **/

        // ejecute en su consola cmd, bash, git los siguientes comandos
        
        // Generar tu clave ssh para poder contribuir al proyecto
        ssh-keygen -t rsa -b 4096 -C "tu email"

        // Configuración global del nombre
        git config --global user.name "John Doe"

        // Configuración global del email
        git config --global user.email johndoe@example.com

        // crea una carpera
        mkdir servicio-app-movil-upla

        // moverse a la carpera
        cd servicio-app-movil-upla
        
        // comando que inicia git
        git init

        // comando que agrega la referencia de la rama
        git remote add origin git@github.com:luissince/servicio-app-movil-upla.git
    
        // comando que descarga los archivos al working directory
        git fetch origin master
        
        // comando que une los cambios al staging area
        git merge origin/master

    2.2 Clonar

        #code

        /** 
        ** Para el proceso de clonación **
        **/

        // Clonar al proyecto
        git clone https://github.com/luissince/servicio-app-movil-upla.git

2. Instale go 

    #Code

        /**
        ** Siga los pasos de instalación de la página oficinal
        **/
        
        // Página oficial
        https://go.dev/

3. Ejecute en la carpte la clonada **go mod download** para descargar las dependencias del proyecto

    #Code

        go mod download

4. Copiar el arhivo de la ruta del EndPoint

    #code

        //copie el archivo .env.example a .env 
        cp .env.example .env

5. Configure la variables de entorno del archivo .env 

    #code

        // Puerto y dirección para levantar el servicio
        GO_PORT="127.0.0.1"

        // Ruta para guardar los archivos logs
        RUTA_LOG='./home/carpera/log.txt'

        // Dirección IP o nombre del servidor
        SERVER_DB=

        // Puerto del servidor
        PORT_DB=
        
        // Nombre de la base de datos del servidor
        NAME_DB=
        
        // Nombre de usuario del servidor
        USER_DB=
        
        // Clave de usuario del servidor
        PASSWORD_DB=

        // Zona horaria del servidor
        TZ_LOCATION="America/Lima"
        

6. Ejecute **go run .** para iniciar ejecutar el servicio   

    #Code

        go run .

7. Cuando se realiza nuevos cambios se tiene que registrar y publicar en la rama correspondiente

    #Code

        // Comando que agrega los cambios realizados en los archivos al área de preparación.
        git add .

        // Comando que agrega los cambios al área de preparación para guardar los cambios en el historial de repositorio.
        git commit -m "Informar cambios"

        // Comando que envía los cambios al repositorio remoto.
        git push origin master
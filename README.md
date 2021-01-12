# Script for Illiux
Para ejecutar el script se deben incorporar los siguientes parámetros
1. Command:    args[0],  // puede ser add o delete
2. Endpoint:   args[1],  // url del server de illiux
3. GoRoutines: routines, // cantidad de routines en paralelo (default 10)
4. OutputPath: args[3],  // path del archivo de errores
5. AuthToken:  args[4],  // token de autenticacion para illiux
6. NCDomain:   args[5],  // Dominio de server de Nextcloud
7. NCUser:     args[6],  // Usuario de server de Nextcloud
8. NCToken:    args[7],  // Basic Token de server de Nextcloud

### Ejemplo: illiux delete https://domain.exmaple.com 5 ~/errors.csv esteeseltokendeilliux https://data.ncexample.com test.account@mailinator.com esteeseltokendeejemploparanextclou


Los dos posibles comandos son:
1. **add**: 
Este comando descarga el archivo de NC y lo lee, por cada elemento leido, se envia un request a illiux para adicionar la subscripcion
2. **delete**: 
Este comando descarga el archivo de NC y lo lee, por cada elemento leido, se envia un request a illiux para eliminar el cliente

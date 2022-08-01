# FuegoQuasar

# Arquitectura implementada

En vista de que los requerimientos dados para la prueba daban libertad para elegir, se implementó el servicio por medio de una aequitectura serverless en la cual las configuraciones y desgastes en detalles de recursos son minimos, ya que el servicio realiza la configuración de forma casi transparente, se debe realizar la generacion de un archivo binario y posteriormente subir a AWS como zip para ser ejecutado por Lambda.

![1_enySPc_XesSQCUWc8i579Qlll](https://user-images.githubusercontent.com/14318998/182242477-163d4a95-b110-40fc-bf53-819622550417.png)

* # Conection: en este package esta implementada la logica que recibe y envia la comunicación con el servicio APIGateway de Lambda AWS
* # Services: en este package estan las dos clases encargadas de los servicios, como son decode que se encarga de decodificar el mensaje y por otra parte la clase      trilateration que implementa el algoritmo de trilateración el cual permite la obtención de la ubicación del portacarga imperial usando la interseccion de los puntos entre circunferencias, basados en la formula matematica propuesta en: https://confluence.slac.stanford.edu/display/IEPM/TULIP+Algorithm+Alternative+Trilateration+Method

![trilateration](https://user-images.githubusercontent.com/14318998/182245790-9063857e-9f5e-4c92-9ea8-04e480b93bc0.png)
Imagen tomada de: https://confluence.slac.stanford.edu/display/IEPM/TULIP+Algorithm+Alternative+Trilateration+Method 

* # Topsecret: este package es el encargado de recibir los parametros del y construir lo que será la respuesta en la clase topsecret, y por ultimo en la clase entities es donde estan almacenados los objetos o estructuras necesarias para almacenar la información a procesar.

Finalmente se tiene el archivo main que solo cuenta con un disparador para iniciar la ejecución de nuestra API REST

# Ejecución

la siguiente url es la asignada por AWS para consumir la API REST 
https://0861uyj7o5.execute-api.us-east-1.amazonaws.com/v1/topsecret

el script propuesto en el ejecicio es el siguiente:
{
    "satellites": [
        {
            "name": "kenobi",
            "distance": 100.0,
            "message": ["este", "", "", "mensaje", ""]
        },
        {
            "name": "skywalker",
            "distance": 115.5,
            "message": ["", "es", "", "", "secreto"]
        },
        {
            "name": "sato",
            "distance": 142.7,
            "message": ["este", "", "un", "", ""]
        }
    ]
    }




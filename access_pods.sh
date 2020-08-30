#!/bin/bash
#Author: Felipe Castillo  -> fcastillovilches@gmail.com

#mostrar namespaces

#kubectl get namespaces


clear

echo ''
echo 'Opciones de Comando'
echo ''

comando=(exec describe logs)

# Agregar Contado
COUNT=1

# Recorrer arreglo de namespaces
for h in "${comando[@]}";
do 
    echo "$COUNT" - $h;    
    comando["$COUNT"]="$h"
    COUNT=$[$COUNT +1]
done

#pedir numero de comando
echo ''
echo 'Introduzca el numero del comando que desea solicitar'
echo ''

#leer id de la opcion seleccionada
read id_comando



clear

# Agregar Contado
COUNTER=1

echo ''
echo 'Opciones de Namespaces'
echo ''

# Recorrer arreglo de namespaces
for i in $(kubectl get namespaces | grep Active |  cut -d ' ' -f1); 
do 
    echo "$COUNTER" - $i;    
    opcion["$COUNTER"]="$i"
    #echo "\$opcion[$COUNTER]" : "${opcion[$COUNTER]}"
    COUNTER=$[$COUNTER +1]
done


#pedir numero de namespaces
echo ''
echo 'Introduzca el numero del namespaces donde se encuentra el pod solicitado'
echo ''

#leer id de la opcion seleccionada
read id_namespaces

clear

# mostar nombre segun id seleccionado
#echo "${opcion[$id_namespaces]}"

# Obtener nombre de namespaces segun ID
kubectl get namespaces | cut -d ' ' -f1 | grep Active

#pedir el dato al usuario
#echo 'Introduzca el namespaces en el que se encuentra el pod'

#leer el dato del teclado y guardarlo en la variable de namespaces

#read namespaces

#mostrar pods del namespaces    
#kubectl get pods -n ${opcion[$id_namespaces]}

echo ''
echo 'Opciones de Namespaces'
echo ''

COUNTER1=1
# Recorrer arreglo de namespaces
for j in $(kubectl get pods -n ${opcion[$id_namespaces]}  |  cut -d ' ' -f1); 
do 
    echo "$COUNTER1" - $j;    
    pod["$COUNTER1"]="$j"
    #echo "\$opcion[$COUNTER]" : "${opcion[$COUNTER]}"
    COUNTER1=$[$COUNTER1 +1]
done

#pedir el dato al usuario
#echo 'Introduzca el pod al que desea entrar'

#leer el dato del teclado y guardarlo en la variable de pod
#read pod

#pedir numero de pod
echo ''
echo 'Introduzca el numero del pod al que desea acceder'
echo ''

#leer id de la opcion seleccionada
read id_pod

clear


#Mostrar el valor de la variable pod y namespaces


echo pod =  ${pod[$id_pod]}
echo namespaces = ${opcion[$id_namespaces]}
echo comando = ${comando[$id_comando]}

echo

clear

if [ $id_comando -eq 1 ]
then
    # Ejecutar comando
    kubectl exec -it pod/${pod[$id_pod]} /bin/bash -n ${opcion[$id_namespaces]}
elif [ $id_comando -eq 2 ]
then

    # Ejecutar comando
    kubectl describe pod/${pod[$id_pod]} -n ${opcion[$id_namespaces]}

else
    # Ejecutar comando
    kubectl logs ${pod[$id_pod]} -n ${opcion[$id_namespaces]}
fi






#Avisar al usuario que se ha terminado de ejecutar el script 

echo ---------Fin del script.-------------

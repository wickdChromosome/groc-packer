#!/bin/bash

set -eou pipefail

if [ $1 == "compile" ]; then

	echo "Compiling project.."
	cd /src/build &&
		make

else

	echo "Running project.."
	/src/build/packing_site

fi	


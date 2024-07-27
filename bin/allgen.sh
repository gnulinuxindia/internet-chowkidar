#!/bin/bash
set -eux

bash bin/apigen.sh &
bash bin/dbgen.sh &

# these three must be in this specific order
# since providergen depends on the mocks from mockgen
# and wire, ran from gogen, depends on the providers from providergen
bash bin/mockgen.sh
bash bin/providergen.sh
bash bin/gogen.sh

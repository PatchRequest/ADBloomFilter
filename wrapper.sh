#!/bin/bash
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

BLUE_CROSS="${BLUE}[+]${NC}"
GREEN_MINUS="${GREEN}[-]${NC}"

printf $BLUE_CROSS" Starting the System...\n"

printf $GREEN_MINUS" Insert Domain: "
read domain

printf $GREEN_MINUS" Insert Target Domain-Controller: "
read dc

printf $GREEN_MINUS" Insert Username: "
read username

printf $GREEN_MINUS" Insert Password(hidden input): "
read -s password
printf "\n" 


printf $BLUE_CROSS" Starting the DCSync Module...\n" 
chmod +x ./syncer/main.py
python3 ./syncer/main.py "$domain/$username:$password@$dc" -outputfile "output" -just-dc-ntlm



printf $BLUE_CROSS" Starting the checker Module...\n"
cd checker
./checker
cd ..
rm output.ntds
printf $BLUE_CROSS" Lets look at the baddies...\n"


printf "${GREEN}+${NC}${BLUE}---------------------------------${NC}${GREEN}+${NC}\n"
cat ./checker/badOnes.txt
printf "${GREEN}+${NC}${BLUE}---------------------------------${NC}${GREEN}+${NC}"

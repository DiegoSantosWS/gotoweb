#Saindo do terminal da DO
$ exit

$ cd nameapp

nameapp$ git init
# resultado do git init


# git remote add dokku dokku@seu-ip-digital-ocean:nameapp
nameapp$ git remote add dokku dokku@seu-ip-digital-ocean:nameapp

# COMANDOS GIT
nameapp$ git add .

nameapp$ git commit -m "Primeiro commit na DO"

nameapp$ git push dokku master
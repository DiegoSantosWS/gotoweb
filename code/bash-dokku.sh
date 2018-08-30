$ dokku apps:create nameapp
#saida
Creating... nameapp

# usando mysql

$ dokku plugin:install https://github.com/dokku/dokku-mysql.git mysql

# Após a instalação
$ dokku mysql:create mydatabase

#Link o database com a aplicação
$ dokku mysql:link mydatabase nameapp

#tem uma saida enorme pois vai ser baixado imagem e gerar um container



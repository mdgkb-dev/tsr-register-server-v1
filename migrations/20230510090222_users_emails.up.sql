update users
set email = concat(login, '@gmail.com')  where users.login !='';
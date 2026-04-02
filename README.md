# E-Biznes

## [Zadanie 1 - Docker](./1/)

- [x] 3.0 [obraz ubuntu z Pythonem w wersji 3.10](./1/3.0.Dockerfile)
- [x] 3.5 [obraz ubuntu:24.0~~2~~4 z Javą w wersji 8 oraz Kotlinem](./1/3.5.Dockerfile)
- [x] 4.0 do powyższego należy [dodać najnowszego Gradle’a oraz paczkę JDBC SQLite](./1/4.0.Dockerfile) w ramach [projektu na Gradle (build.gradle)](./1/build.gradle)
- [x] 4.5 stworzyć [przykład typu HelloWorld](./1/HelloWorld.kt) oraz [uruchomienie aplikacji przez CMD oraz gradle](./1/4.5.Dockerfile)
- [x] 5.0 dodać [konfigurację docker-compose](./1/docker-compose.yaml)

Termin: 25.03

Punkty 3.0-4.5 powinny mieć osobny obraz Dockerowy.

Obraz dockerowy należy wrzucić na hub.docker.com.
Dockerfile oraz dodatkowe pliki powinny być na repozytorium git.
Readme powinno zawierać [link do obrazu na hub.docker.com](https://hub.docker.com/r/jmarkiewicz0/ebiznes-1).

## [Zadanie 2 - Scala](./2/)

Należy stworzyć aplikację na frameworku ~~Play~~ lub Scalatra.

- [x] 3.0 Należy stworzyć [kontroler do Produktów](./2/src/main/scala/example/ebiznes/ebiz2/Products.scala)
- [x] 3.5 Do kontrolera należy stworzyć endpointy zgodnie z CRUD - dane pobierane z listy
- [x] 4.0 Należy stworzyć kontrolery do [Kategorii](./2/src/main/scala/example/ebiznes/ebiz2/Categories.scala) oraz [Koszyka](./2/src/main/scala/example/ebiznes/ebiz2/Cart.scala) + endpointy zgodnie z CRUD
- [x] 4.5 Należy aplikację uruchomić na dockerze ([stworzyć obraz](./2/Dockerfile)) oraz dodać [skrypt uruchamiający aplikację via ngrok](./2/run.sh)
- [x] 5.0 Należy dodać [konfigurację CORS](./2/src/main/scala/ScalatraBootstrap.scala) dla dwóch hostów dla metod CRUD

Kontrolery mogą bazować na listach zamiast baz danych. CRUD: show all, show by id (get), update (put), delete (delete), add (post).

<https://scalatra.org/getting-started/first-project.html>
~~<https://www.playframework.com/>~~

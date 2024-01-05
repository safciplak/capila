# Documentation of the Capila Library
  
This package serves as a library for API development.

Be aware that this is still a package under development. Some things can change in the near future.

Please look at the flow diagram in this directory.  
you can open this by going to the following link:  
https://www.draw.io/
(or download the drawio application)

## Explanation of the total flow

There's a thought behind this flow.

> Bootstrap
The external application (your customer's api) starts and makes the endpoints known to the webserver.

> Web
The webserver will call the correct external handler's endpoint.  
(You can shield each endpoint with a JWT authentication.)

> External handler
An external handler has one or more endpoints. Each endpoint refers to a function within that handler. The function uses a presenter to create the output-format of the data provided by the corresponding service.

A handler also feeds the underlying service with the correct parameters.

Since handlers are (relatively) slow to test, they are very thin and act as a pass-through.  (You can skip slow tests with the make test-s option.)
Advantage is, that most of the work is done from the service-layer and deeper which makes testing of a service very fast, because there is no need for a webserver.

> External service
An external service is the worker for the external handler.

A service can validate the input given by the handler and call some business logic that might be required.

If a service needs data from some data source, it calls a corresponding repository.

> External repository
An external repository is an intermediair between actual data and the application. It will communicate with a data source and takes care of getting data inside the application or sending it to the data source. (Data can come from external sources, like a database, a text file, an Excel sheet or even another API.)

A repository will get specific domain models and give them back to the calling service in a common format.

> External model
An external model is a set of fields combined into an object (struct).


## Composition of this package

This package is a library and can be used from a customer's API application.

You can checkout the blue-api-demo repository to see how you could use the library.

## Uitleg bouwtekening

In het business-deel wordt de daadwerkelijke code geschreven waarvoor de applicatie bedoeld is. In principe bestaat het business-deel uit een service die dus het werk doet voor een handler. Dit is ook de plek waar de business logica wordt toegepast.

Als de service data nodig heeft, heeft deze geen weet van de plek/manier waar de data vandaan komt en zal deze data opvragen via een repository. De repository weet als enige waar specifieke data vandaan komt. Dat betekent dat alleen een repository moet worden aangepast als de locatie wijzigt of het type database of ...
(Data zou uit een excel/csv/redis/mysql/postgres/config bestand kunnen komen.)
De repository kent de onderliggende data structuur aan de hand van modellen.

## directory initial : initialize

Hier staat de initialisatie van de applicatie.

## directory web : webserver / router / middleware

Hier staat alles dat de communicatie met het web verzorgd.

> Webserver

- is redelijk standaard gemaakt.
- sluit netjes af nadat het laatste request is beantwoord.
- heeft maximale timeouts

> Router

- roept benodigde handlers aan om in detail de routing toe te voegen
- dit zorgt ervoor dat de daadwerkelijke router configuratie daar gedaan wordt,
	waar de wijziging ook wordt aangebracht

> Middelware

Deze werkt als een schil om alle requests heen en bestaat uit verschillende onderdelen:
-  info
	- drukt de url en de hoeveelheid tijd af
- isAuthorized
	- controleert of er autorisatie nodig is
	- wordt bepaald boven in bij skipAuth
- recovery
	- handelt onverwachte fouten af

## directory handling

Bestaat uit 2 subdirectories die te samen de afhandeling van het web-request doen.

	### directory handlers

	Hier staat alles dat de afhandeling van iedere endpoint verzorgt.

	Een handler is een onderdeel van de applicatie dat nog bij het "netwerk" hoort.
    - deze praat nog in request/response
    - bepaalt hoe de data gepresenteerd moet worden door de desbetreffende presenter te gebruiken.

	AddRoutes()
    - deze wordt aangeroepen door de web/router
    - hierin staan de endpoints die deze handler afhandelt
    - let op dat ieder endpoint over de gehele applicatie uniek moet zijn

	> HomeHandler

	Als je bijvoorbeeld kijkt naar de HomeHandler, dan zie je bij AddRoutes een verwijzing naar handler.Info en deze Info-functie staat daar onder. 

	Zou je dus bijvoorbeeld "About" functionaliteit bij HomeHandler willen toevoegen, dan copier je de Info functie, noemt deze About en doet daarin de benodigde aanpassingen. Voeg deze functie toe aan de eerder genoemde AddRoutes en dan is het "aan elkaar geknoopt".

	> UserHandler

	Deze handler is veel meer een handler volgens de bouwtekening, want deze laat het daadwerkelijke werk over aan de service-laag. Zo zie je dat /users dus naar de functie List() gaat maar dat deze functie zelf aan de service-laag om de daadwerkelijke data vraagt om de lijst te vullen. Dit zorgt dus voor een goede scheiding, waardoor testen eenvoudiger wordt.


	### directory presenters

	Hier staat alles dat de output verzorgt die door handlers benodigd is.

## directory business

Bestaat uit 3 subdirectories die te samen de business en data bevatten van de applicatie.

	### directory services

	Hier staan de services die het werk voor de handlers doen. Een service kan zelfstandig dingen doen en dus zelfstandig getest worden.

	Een service kan ook met data willen werken die "ergens" vandaan moet komen. In dat geval zal de service een repository aanroepen die deze data gaat verzorgen of wegschrijven/wissen, noem maar op.

	Voor een handler is de service dus het object dat de opdracht uitvoert. De service weet niet wat de handler er mee wil en voert echt alleen het werk uit en zo kan een List() functie binnen een service kan gebruikt worden om uiteindelijk te eindigen als JSON, maar zou even zo goed kunnen eindigen als SOAP-XML of ... Dit is afhankelijk van de presenter die door de handler gebruikt wordt en dus voor het functioneren van het business-deel totaal onbelangrijk.

	### directory repositories

	Een repository bepaalt waar de specifieke data vandaan komt. Voor de rest van de applicatie is dit verder niet van belang zolang de repository het gevraagde maar kan leveren.  

	### directory models

	Een model is de representatie is van een stuk data.

Deze blauwdruk-applicatie werkt momenteel met een lokale database (sqlite3) en een orm (gorm) waardoor er vrij eenvoudig extra zaken zijn te realiseren.


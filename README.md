# Registre des ventes

Ce projet à pour but de créer une interface REST qui permet de gérer un registre des ventes de biens. Il y a donc une notion de client, de type de produit, et de vente.

## Routes

### Publiques : 
- `POST /login {email, password}`

### Privées :

Users :
- `GET /users`
- `GET /users/:id`
- `POST /users {first_name, last_name, email, password}`
- `PATCH /users/:id {first_name?, last_name?, email?, password?}`
- `DELETE /users/:id`

Clients :
- `GET /clients`
- `GET /clients/:id`
- `POST /clients {first_name, last_name, birth_date, birth_place, id_card_path, id_card_type, id_card_num, id_card_authority, id_card_date}`
- `PATCH /clients/:id {first_name?, last_name?, birth_date?, birth_place?, id_card_path?, id_card_type?, id_card_num?, id_card_authority?, id_card_date?}`
- `PATCH /clients/:id/validate`
- `PATCH /clients/:id/unvalidate`
- `DELETE /clients/:id`

Produits :
- `GET /products`
- `GET /products/:id`
- `POST /products {name, pic, agreement, active_weight, category}`
- `PATCH /products {name?, pic?, agreement?, active_weight?, category?}`
- `DELETE /products/:id`

Ventes :
- `GET /sells`
- `GET /sells/:id`
- `POST /sells {payment_method, products, client_id}`


## Modeles 

Client
- ID
- Validé
- Document identité
- Type document identité
- Numéro document identité
- Autorité de délivrance
- Date de délivrance
- Nom
- Prénom
- Date de naissance
- Lieu de naissance

Produit
- ID
- Dénomination
- Photo
- Agrément
- Matière active
- Catégorie

Vente
- ID
- Date et heure d'achat
- Mode de paiement
- Produits Achetés
- Client
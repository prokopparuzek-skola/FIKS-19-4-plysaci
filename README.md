---
header-includes:
	\usepackage[czech]{babel}
	\usepackage{a4wide}
---
# Plyšáci

## Řešení

Úlohu řeším metodou backtrackingu. Zkouším na volná místa dosazovat spadlé poníky, když to vyjde pokračuji dál, když ne 
vrátím se k předchozí úrovni kdy to ještě šlo. (Prohedávám do hloubky graf možných řešení)

## Složitost

V nejhorším možnám případě kdybych prohledal všechny možnosti by složitost byla exponenciální, protože graf řešení roste 
exponenciálně, každý vrchol má tolik potomků kolik ještě zbývá poníků na zemi.

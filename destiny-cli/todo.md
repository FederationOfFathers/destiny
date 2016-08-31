
x acct summary {user}                         /{membershipType}/Account/{destinyMembershipId}/Summary/
  acct advisors {user}                        /{membershipType}/Account/{destinyMembershipId}/Advisors/
  acct items {user}                           /{membershipType}/Account/{destinyMembershipId}/Items/
  acct grimoire {user}                        /Vanguard/Grimoire/{membershipType}/{membershipId}/
x acct stats {user}                           /Stats/Account/{membershipType}/{destinyMembershipId}/
  acct triumphs {user}                        /{membershipType}/Account/{destinyMembershipId}/Triumphs/

  char activityhistory {user} {character}     /Stats/ActivityHistory/{membershipType}/{destinyMembershipId}/{characterId}/
  char activities {user} {character}          /{membershipType}/Account/{destinyMembershipId}/Character/{characterId}/Activities/
  char advisorv2 {user} {character}           /{membershipType}/Account/{destinyMembershipId}/Character/{characterId}/Advisors/V2/
  char inventory data {user} {character}      /{membershipType}/Account/{destinyMembershipId}/Character/{characterId}/Inventory/
  char inventory summary {user} character}    /{membershipType}/Account/{destinyMembershipId}/Character/{characterId}/Inventory/Summary/
  char progression {user} {character}         /{membershipType}/Account/{destinyMembershipId}/Character/{characterId}/Progression/
  char summary {user} {character}             /{membershipType}/Account/{destinyMembershipId}/Character/{characterId}/
x char agstats {user} {character}             /Stats/AggregateActivityStats/{membershipType}/{destinyMembershipId}/{characterId}/
x char stats {user} {character}               /Stats/{membershipType}/{destinyMembershipId}/{characterId}/
  char item get {user} {character} {item}     /{membershipType}/Account/{destinyMembershipId}/Character/{characterId}/Inventory/{itemInstanceId}/
  char item ref {user} {character} {hash}     /{membershipType}/Account/{destinyMembershipId}/Character/{characterId}/ItemReference/{itemHash}/
  char exotic stats {user} {character}        /Stats/UniqueWeapons/{membershipType}/{destinyMembershipId}/{characterId}/

  activity {id}                               /Stats/PostGameCarnageReport/{activityId}/

  search {name}                               /SearchDestinyPlayer/{membershipType}/{displayName}/

  define advisors                             /Advisors/V2/
  define manifest data                        /Manifest
x define manifest download
  define item {type} {id}                     /Manifest/{type}/{id}/
  define grimoire                             /Vanguard/Grimoire/Definition/
  define stats                                /Stats/Definition/



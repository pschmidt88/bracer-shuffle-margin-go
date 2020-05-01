# Technical Flow

## Connected Realm

- find connected realm id by realm name

using Realm API

`GET /data/wow/realm/{realmSlug}`

## Get Auctions

- Get all auctions

using Action House API

`GET /data/wow/connected-realm/{connectedRealmId}/auctions`

## Filter auctions

We need prices for a bunch of items

### Vendor Materials

- 159959 [Nylon Thread](https://www.wowhead.com/item=159959/nylon-thread) 
    > Fixed price 60s - vendor mat. Sometimes this is cheaper on the auction house, so maybe in a later version we can search for it in the auction house too.

### Raw Materials

- 152576 [Tidespray Linen](https://www.wowhead.com/item=152576/tidespray-linen)
- 154722 [Tempest Hide](https://www.wowhead.com/item=154722/tempest-hide)
- 154165 [Calcified Bone](https://www.wowhead.com/item=154165/calcified-bone)
- 152541 [Coarsed Leather](https://www.wowhead.com/item=152541/coarse-leather)
- 153050 [Shimmerscale](https://www.wowhead.com/item=153050/shimmerscale)
- 154164 [Blood-stained Bone](https://www.wowhead.com/item=154164/blood-stained-bone)
i
- 153051 [Mistscale](https://www.wowhead.com/item=153051/mistscale)
- 152577 [Deep Sea Satin](https://www.wowhead.com/item=152577/deep-sea-satin)

### Enchanting Materials

- 152876 [Umbra Shard](https://www.wowhead.com/item=152876/umbra-shard)
- 152875 [Gloom Dust](https://www.wowhead.com/item=152875/gloom-dust)
- 152877 [Veiled Crystal](https://www.wowhead.com/item=152877/veiled-crystal)

### Ring Enchants

- 168447 [Accord of Haste](https://www.wowhead.com/item=168447/enchant-ring-accord-of-haste)
- 168446 [Accord of Critical Strike](https://www.wowhead.com/item=168446/enchant-ring-accord-of-critical-strike)
- 168448 [Accord of Mastery](https://www.wowhead.com/item=168448/enchant-ring-accord-of-mastery)
- 168449 [Accord of Versatility](https://www.wowhead.com/item=168449/enchant-ring-accord-of-versatility)

### Weapon Enchants

- 168596 [Force Multiplier](https://www.wowhead.com/item=168596/enchant-weapon-force-multiplier)
- 168593 [Machinist's Brilliance](https://www.wowhead.com/item=168593/enchant-weapon-machinists-brilliance)
- 168598 [Naga Hide](https://www.wowhead.com/item=168598/enchant-weapon-naga-hide)
- 168592 [Oceanic Restoration](https://www.wowhead.com/item=168592/enchant-weapon-oceanic-restoration)


## Calculate Cheapest Bracers

### Tidespray Bracer (Tailor)

`craftingCost = 10*Tidespray Linen + 5*Nylon Thread`

### Coarse Leather Armguards (Leatherworking)

`craftingCost = 6*Coarse Leather + 4*Blood Stained Bones`

### Shimmerscale Armgards

`craftingCost = 6*Shimmerscale + 4*Blood Stained Bones`




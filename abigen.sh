#/bin/sh

names=("erc20" "ritsu" "izumi" "iziPool" "drips" "symmetric" "balancer_vault" "balancer_token" "robinos" "loopring_lock" "dorahacks_voting" "avalon_claim" "loopring_exchangev3" "okidori_marketplace")

for (( i = 0; i < ${#names[@]}; ++i ));
do
    lower=$(echo "${names[i]}" | tr '[:upper:]' '[:lower:]')
    mkdir -p adapters/contracts/$lower
    abigen --abi ./abi/${names[i]}.json \
    --pkg $lower \
    --type ${names[i]} \
    --out adapters/contracts/$lower/${names[i]}.go
done

exit 0

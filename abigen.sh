#/bin/sh

names=("erc20", "ritsu")


for (( i = 0; i < ${#names[@]}; ++i ));
do
    lower=$(echo "${names[i]}" | tr '[:upper:]' '[:lower:]')
    abigen --abi ${names[i]}.json \
    --pkg $lower \
    --type ${names[i]} \
    --out adapters/contracts/$lower/${names[i]}.go
done

exit 0

#/bin/sh

names=("erc20" "ritsu" "izumi" "iziPool")

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

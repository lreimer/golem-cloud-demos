# Golem Cloud Demos

Demo repository for different Golem Cloud WASI applications in different languages.

## Rust Shopping Cart Demo

```bash
# make sure you have setup your Rust environment as described in 
# - https://www.golem.cloud/learn/common-tooling
# - https://www.golem.cloud/learn/rust

golem new --template rust-shopping-cart --component-name shopping-cart
cd shopping-cart

# build and install WASI component
cargo component build --release
cd target/wasm32-wasi/release/
golem component add --component-name shopping-cart lib.wasm
golem instance add --instance-name shopping-cart-1 --component-name shopping-cart

# invoke the instance and its API functions
golem instance invoke-and-await \
    --component-name=shopping-cart \
    --instance-name=shopping-cart-1 \
    --function=golem:component/api/initialize-cart \
    --parameters='["test-user"]'

golem instance invoke-and-await \
    --component-name=shopping-cart \
    --instance-name=shopping-cart-1 \
    --function=golem:component/api/get-cart-contents \
    --parameters='[]'

golem instance invoke-and-await \
    --component-name=shopping-cart \
    --instance-name=shopping-cart-1 \
    --function=golem:component/api/add-item \
    --parameters='[{"product-id": "1234abc", "name": "test-item", "price": 9.99, "quantity": 42}]'

golem instance invoke-and-await \
    --component-name=shopping-cart \
    --instance-name=shopping-cart-1 \
    --function=golem:component/api/get-cart-contents \
    --parameters='[]'
```

## Go HTTP Demo

```bash
# make sure you have setup your Tiny Go environment as described in 
# - https://www.golem.cloud/learn/common-tooling
# - https://www.golem.cloud/learn/go

golem new --template go-actor-full --component-name go-golem-demo
cd go-golem-demo

# build and install WASI component
make build
golem component add --component-name go-golem-demo go_golem_demo.wasm
golem instance add --instance-name go-golem-demo-1 --component-name go-golem-demo

# repeat the following two commands to add numbers and get results
golem instance invoke-and-await \
    --component-name=go-golem-demo \
    --instance-name=go-golem-demo-1 \
    --function=golem:component/api/add \
    --parameters='[42]'

golem instance invoke-and-await \
    --component-name=go-golem-demo \
    --instance-name=go-golem-demo-1 \
    --function=golem:component/api/get \
    --parameters='[]'

```

## Maintainer
 
M.-Leander Reimer (@lreimer), <mario-leander.reimer@qaware.de>

## License

This software is provided under the MIT open source license, read
the `LICENSE` file for details.

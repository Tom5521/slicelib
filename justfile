test:
    go test -v ./tests/*
gh-release tag:
    git tag {{tag}}
    git push --tags
    gh release create {{tag}} --generate-notes

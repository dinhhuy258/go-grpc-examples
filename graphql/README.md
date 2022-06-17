# Graphql

## Built With

- Go
- [gqlgen](https://github.com/99designs/gqlgen)

## How to run

**Initialize graphql project**

```
go run -mod=mod github.com/99designs/gqlgen init
```

**Generate graphql**

```
go run -mod=mod github.com/99designs/gqlgen generate
```

**Start server**

```
go run main.go
```

**Queries**

```
{
  books {
    id,
    name,
    author {
      id,
      name
    }
  }
}
```

```
{
  author(id: 1){
    id,
    name,
    books {
      id,
      name
    }
  }
}
```

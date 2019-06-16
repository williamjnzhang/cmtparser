# golang ComMenT PARSER

base on origin go/parser in golang version go1.11

The origin go/parser won't parse the Comments to the corresponding field of some type of ast node, such as 
+ `ast.Field` in `ast.FuncType -> Params`
+ `ast.Field` in `ast.FuncType -> Results`
+ `ast.Field` in `ast.StructType -> Fields`
+ ...

This modified parser will parse the comments in some structure (listed below) of certain syntax (defined below) into the AST.

Comments form is still align to [origin go comments](https://golang.org/ref/spec#Comments)

Syntax notation is align to [origin](https://golang.org/ref/spec#Notation)

## Function Types

With regard to [function types](https://golang.org/ref/spec#Function_types), we extend `ParameterDecl` to

```
ParameterDecl = [IdentifierList] ["..."] Type .
IdentifierList = {IdentComment} identifier {"," {IdentComment} identifier} .
```
<!--`IdentifierList = identifier {IdentComment} {"," identifier} .`-->

In AST, the `IdentComment` will append to corresponding `ast.Ident -> Comment`.

```
func f1 (/*f1a*/ a string, /*f1b*/ b string) {

}

func f2 (
    /*f3s1*/ s1 string,
    /*f3t2*/ t2 struct {
        i1 int
    },
) {}
```

## Struct Types
Similar to function type, `IdentifierList` of struct type have same syntax to function type, thus `IdentComment` defined above is also valid in struct type.

```
// struct1
type s1 struct {
    /*s1c1*/ a string
    // s1c3
    c string
}

// struct2
type s2 struct {/*s2c2*/b string; /*s2c3*/ c string}
```
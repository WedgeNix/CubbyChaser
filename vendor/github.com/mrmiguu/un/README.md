# un
un.Wrap your errors

<i>Before:</i>
```go
b, err := json.Marshal(data)
if err != nil {
    panic(err)
}

...

reData := Data{}
err = json.Unmarshal(b, &reData)
if err != nil {
    panic(err)
}
```

<i>After!</i>
```go
b := un.Bytes(json.Marshal(data))

...

reData := Data{}
un.Wrap(json.Unmarshal(b, &reData))
```

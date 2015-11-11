## chronos-client

[![Build Status][travis-image]][travis-url] [![Coverage][coverage-image]][coverage-url] [![GoDoc][godoc-image]][godoc-url] [![Release][release-image]][release-url]

A Go Chronos client library that provides Chronos operations.

### Installation

```
go get github.com/yieldbot/chronos-client
```

### Usage

#### Getting jobs

See [jobs.go](examples/jobs/jobs.go) for full code.

```go
jobs, err := chronos.Client.Jobs()
if err != nil {
  log.Fatal(err)
}
for _, j := range jobs {
  fmt.Printf("%s\n", j.Name)
}
```

#### Adding a job

See [add.go](examples/add/add.go) for full code.

```go
var j = `{"schedule": "R/2015-11-09T00:00:00Z/PT24H", "name": "test-1", "epsilon": "PT30M", "command": "echo test1 && sleep 60", "owner": "localhost@localhsot", "async": false}`
_, err := chronos.Client.AddJob(j)
if err != nil {
  log.Fatal(err)
}
fmt.Printf("The job is added\n")
```

#### Running a job

See [run.go](examples/run/run.go) for full code.

```go
_, err := chronos.Client.RunJob("test-1", "")
if err != nil {
  log.Fatal(err)
}
fmt.Printf("test-1 job is running\n")
```

#### Killing job tasks

See [killtasks.go](examples/killtasks/killtasks.go) for full code.

```go
_, err := chronos.Client.KillJobTasks("test-1")
if err != nil {
  log.Fatal(err)
}
fmt.Printf("test-1 job tasks are killed\n")
```

#### Deleting a job

See [delete.go](examples/delete/delete.go) for full code.

```go
_, err := chronos.Client.DeleteJob("test-1")
if err != nil {
  log.Fatal(err)
}
fmt.Printf("test-1 job is deleted\n")
```

### License

Licensed under The MIT License (MIT)  
For the full copyright and license information, please view the LICENSE.txt file.

[travis-url]: https://travis-ci.org/yieldbot/chronos-client
[travis-image]: https://travis-ci.org/yieldbot/chronos-client.svg?branch=master

[godoc-url]: https://godoc.org/github.com/yieldbot/chronos-client
[godoc-image]: https://godoc.org/github.com/yieldbot/chronos-client?status.svg

[release-url]: https://github.com/yieldbot/chronos-client/releases/tag/v1.0.0
[release-image]: https://img.shields.io/badge/release-v1.0.0-blue.svg

[coverage-url]: https://coveralls.io/github/yieldbot/chronos-client?branch=master
[coverage-image]: https://coveralls.io/repos/yieldbot/chronos-client/badge.svg?branch=master&service=github)
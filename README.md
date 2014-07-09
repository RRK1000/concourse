# concourse

TODO: What is it?

### Roadmap

[Tracker](https://www.pivotaltracker.com/n/projects/1059262)

### Try it on Vagrant

1. Install dependencies

```
vagrant plugin install vagrant-bosh
gem install bosh_cli --no-ri --no-rdoc
go get github.com/concourse/fly
```

1. Create a new VM

```
vagrant up
```

1. Play around with [ATC](https://github.com/concourse/atc)
  - Browse to your [local ATC](http://127.0.0.1:8080) and trigger a build or two.
  - Edit `manifests/vagrant-bosh.yml` and `vagrant provision` to reconfigure your builds.

1. Play around with [Fly](https://github.com/concourse/fly)
  - Write a build config (`build.yml`) and run it with `fly`. See [Turbine's](https://github.com/concourse/turbine/blob/master/build.yml) for an example.

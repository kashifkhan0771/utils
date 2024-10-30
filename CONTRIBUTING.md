# Contributing to utils

This is a short guide on how to contribute to the project.

## Submitting a pull request

If you find a bug that you'd like to fix, or a new feature that you'd like to implement then please submit a pull request via GitHub.


Fork the Repository:
1. Visit https://github.com/kashifkhan0771/utils
2. Click the "Fork" button to create your own fork
3. Clone your fork locally:

    git clone git@github.com:<your-username>/utils.git
    cd utils

Make a branch to add your new feature

    git checkout -b my-new-feature main

And get hacking.

When ready - run the unit tests for the code you changed

    make test

Make sure you

* Add documentation for a new feature
* Add unit tests for a new feature
* squash commits down to one per feature
* rebase to develop `git rebase main`

When you are done with that

    git push origin my-new-feature

Your patch will get reviewed, and you might get asked to fix some stuff.

If so, then make the changes in the same branch, squash the commits, rebase it to develop then push it to GitHub with `--force`.

## Test

Tests are run using a testing framework, so at the top level you can run this to run all the tests.

```bash
# runs all tests
make test
```

## Adding New Dependency

```bash
RUNTHIS='go get <package>'
```

#### Example

```bash
RUNTHIS='go get github.com/sirupsen/logrus'
```

```bash
RUNTHIS='go get github.com/sirupsen/logrus@1.7.0'
```

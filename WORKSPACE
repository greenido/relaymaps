git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    commit = "bfa3601d9ab664b448ddb4cc7e48eea511217aaf",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories", "new_go_repository")
go_repositories()

# go-polyline to generate Google Maps polylines.
new_go_repository(
    name = "com_github_twpayne_gopolyline",
    importpath = "github.com/twpayne/go-polyline",
    commit = "6431f8c69af3095d8d3c9a8a227975981452a9c9",
)

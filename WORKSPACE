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

new_go_repository(
    name = "com_github_wcharczuk_gochart",
    importpath = "github.com/wcharczuk/go-chart",
    commit = "66b99eb8e38fd849fe4a78262d0fac8fcd46dc71",
)

new_go_repository(
    name = "com_github_vincentpetithory_dataurl",
    importpath = "github.com/vincent-petithory/dataurl",
    commit = "9a301d65acbb728fcc3ace14f45f511a4cfeea9c",
)

# Those are dependencies, mostly from gochart.
new_go_repository(
    name = "com_github_golang_freetype",
    importpath = "github.com/golang/freetype",
    commit = "d9be45aaf7452cc30c0ceb1b1bf7efe1d17b7c87",
)

new_go_repository(
    name = "org_golang_x_image",
    importpath = "golang.org/x/image",
    commit = "069db1da13841f750892417d97b1b059aec960cd",
)

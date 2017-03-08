load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "main",
    srcs = [
        "elevation.go",
        "load.go",
        "main.go",
        "pages.go",
        "print.go",
        "structs.go",
    ],
    deps = [
      "@com_github_twpayne_gopolyline//:go_default_library",
      "@com_github_vincentpetithory_dataurl//:go_default_library",
      "@com_github_wcharczuk_gochart//:go_default_library",
      "@com_github_wcharczuk_gochart//drawing:go_default_library",
    ],
)

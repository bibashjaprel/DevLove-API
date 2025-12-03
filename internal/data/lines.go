package data

import "math/rand"

var GolangLines = []string{
    "Are you a goroutine? Because my heart keeps running every time I think of you.",
    "Are you a pointer? Because you always reference my heart.",
    "Are you Go modules? Because you bring order to my life.",
    "You must be a channel, because you synchronize my feelings.",
}

var GitLines = []string{
    "Are you GitHub? Because I want to commit my life to you.",
    "Are you a merge request? Because I want to resolve all our conflicts.",
    "Are you a branch? Because I always want to check you out.",
    "Are you a pull request? Because you bring new features to my life.",
}

var DockerLines = []string{
    "Are you a Docker image? Because my heart pulls you every time.",
    "Are you a container? Because you encapsulate all my love.",
    "Are you a Dockerfile? Because you build my happiness step by step.",
    "Are you a volume? Because I want to persist you forever.",
}

var KubernetesLines = []string{
    "Are you Kubernetes? Because you orchestrate my whole world.",
    "Are you a pod? Because you run in my heart.",
    "Are you a service? Because you expose my feelings.",
    "Are you a deployment? Because you keep my love rolling updates.",
}

var RomanticLines = []string{
    "Are you a 200 OK? Because everything feels right when I see you.",
    "Are you an API key? Because I should never share you with anyone.",
    "Are you a database? Because my heart queries you every day.",
    "Are you a front-end? Because you make my world beautiful.",
    "Are you a backend? Because you support me in every way.",
}

func GetRandomLine(category string) (string, string) {
    var lines []string
    var cat string
    switch category {
    case "golang":
        lines = GolangLines
        cat = "golang"
    case "git":
        lines = GitLines
        cat = "git"
    case "docker":
        lines = DockerLines
        cat = "docker"
    case "kubernetes":
        lines = KubernetesLines
        cat = "kubernetes"
    case "romantic":
        lines = RomanticLines
        cat = "romantic"
    case "random":
        all := append(GolangLines, GitLines...)
        all = append(all, DockerLines...)
        all = append(all, KubernetesLines...)
        all = append(all, RomanticLines...)
        lines = all
        cat = "random"
    default:
        lines = RomanticLines
        cat = "romantic"
    }
    if len(lines) == 0 {
        return cat, "No pickup lines found."
    }
    return cat, lines[rand.Intn(len(lines))]
}

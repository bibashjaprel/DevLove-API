package data

import "math/rand"

var GolangLines = []string{
    "Are you a goroutine? Because my heart starts racing whenever you're around.",
    "Are you a pointer? Because you always point me in the right direction.",
    "Are you Go modules? Because you bring stability to my chaotic life.",
    "Are you a channel? Because you communicate directly with my soul.",
    "Are you defer? Because I’d love to hold you until the very end.",
    "Are you Go build? Because you compile perfectly with my feelings.",
    "Is your name fmt? Because you format my life beautifully.",
    "Are you a slice? Because my love for you keeps growing dynamically.",
    "Are you an interface? Because you define everything I ever wanted.",
    "Are you a map? Because I keep looking up my happiness in you.",
}

var GitLines = []string{
    "Are you GitHub? Because I can't stop starring you.",
    "Are you a commit? Because I want to keep you in my history forever.",
    "Are you a merge request? Because I want to resolve all our conflicts.",
    "Are you a branch? Because I always want to check you out.",
    "Are you a pull request? Because you bring new changes into my life.",
    "Are you a rebase? Because you make my past look cleaner.",
    "Are you Git stash? Because I keep hiding my feelings for you.",
    "Are you a version tag? Because you're a stable release.",
    "Are you Git init? Because you just started something in my heart.",
    "Are you a fork? Because I want to explore your world.",
}

var DockerLines = []string{
    "Are you a Docker image? Because I can't help pulling you.",
    "Are you a container? Because you keep my feelings isolated and safe.",
    "Are you Docker build? Because you assemble my joy step by step.",
    "Are you a volume? Because my love persists even after restart.",
    "Are you a Docker compose? Because everything works better when we’re together.",
    "Are you a port? Because I want to expose my heart to you.",
    "Are you a health check? Because you keep me stable.",
    "Are you a base image? Because everything good starts from you.",
    "Are you a tag? Because you give meaning to everything.",
    "Are you Docker run? Because my heart starts instantly when I see you.",
}

var KubernetesLines = []string{
    "Are you Kubernetes? Because you orchestrate my world effortlessly.",
    "Are you a pod? Because you run inside my heart.",
    "Are you a service? Because you expose my hidden feelings.",
    "Are you a deployment? Because my love scales with you.",
    "Are you a configmap? Because you store all my secrets.",
    "Are you a secret? Because I keep you safe and encrypted.",
    "Are you an ingress? Because you route straight into my heart.",
    "Are you a node? Because I schedule all my dreams on you.",
    "Are you HPA? Because my love autoscale whenever I see you.",
    "Are you kubectl? Because with you everything becomes easier.",
}

var RomanticLines = []string{
    "Are you a 200 OK? Because everything feels right with you.",
    "Are you an API key? Because I never want to share you with anyone else.",
    "Are you a database? Because my heart queries you every moment.",
    "Are you a frontend? Because you beautify everything in my life.",
    "Are you a backend? Because you support me silently.",
    "Are you an SSL certificate? Because you secure all my connections.",
    "Are you a heartbeat signal? Because you keep me alive.",
    "Are you a UI? Because you’re simply stunning.",
    "Are you sunshine? Because you light up everything.",
    "Are you bandwidth? Because my heart speeds up around you.",
}

var LinuxLines = []string{
    "Are you Arch Linux? Because I’d brag about you everywhere.",
    "Are you sudo? Because you give me the permission to do anything.",
    "Are you a shell? Because my commands always depend on you.",
    "Are you systemd? Because you start everything in my life.",
    "Are you grep? Because I find everything I look for in you.",
    "Are you a kernel? Because you handle all my processes.",
    "Are you chmod 777? Because you give full access to my heart.",
    "Are you a symlink? Because I’m always connected to you.",
}

var NetworkingLines = []string{
    "Are you a router? Because you direct my heart.",
    "Are you an IP address? Because I can’t connect without you.",
    "Are you DNS? Because you resolve my confusion.",
    "Are you bandwidth? Because everything speeds up with you.",
    "Are you a VPN? Because you make me feel safe.",
    "Are you latency? Because you keep me waiting… but it’s worth it.",
}

var DatabaseLines = []string{
    "Are you SQL? Because my queries return only you.",
    "Are you a primary key? Because there’s no one like you.",
    "Are you a foreign key? Because we’re linked forever.",
    "Are you ACID? Because you make my love consistent.",
    "Are you replication? Because my love copies everywhere.",
    "Are you an index? Because you make everything faster.",
}

var CloudLines = []string{
    "Are you AWS? Because you’re the cloud of my dreams.",
    "Are you S3? Because I want to store all my memories with you.",
    "Are you EC2? Because you compute my emotions.",
    "Are you CloudFront? Because you deliver my happiness globally.",
    "Are you IAM? Because you give me the right permissions.",
}

var AILines = []string{
    "Are you a neural network? Because you learn everything about me.",
    "Are you GPT? Because you complete my sentences.",
    "Are you training data? Because I can't grow without you.",
    "Are you fine-tuning? Because you improve my every day.",
    "Are you reinforcement learning? Because I get rewarded when I choose you.",
}

var SecurityLines = []string{
    "Are you HTTPS? Because you encrypt my heart.",
    "Are you firewall? Because you protect me from the world.",
    "Are you 2FA? Because I feel safer with you.",
    "Are you zero trust? Because you’re the only one I trust fully.",
}

var DevOpsLines = []string{
    "Are you CI/CD? Because you automate my happiness.",
    "Are you Terraform? Because my world builds around you.",
    "Are you Ansible? Because you configure me perfectly.",
    "Are you a pipeline? Because I flow better with you.",
    "Are you logs? Because I keep checking up on you.",
}

var NepaliLines = []string{
    "Are you momo? Because I can never say no to you.",
    "Are you Fewa Lake? Because I get lost in your beauty.",
    "Are you Nepali tea? Because you warm my heart.",
    "Are you Dashain? Because you bring joy every year.",
    "Are you KTM traffic? Because you drive me crazy… but I still stay.",
    "Are you a bandh? Because you stop everything when you're around.",
}

func GetRandomLine(category string) (string, string) {
    var lines []string
    var cat string

    switch category {
    case "golang":
        lines, cat = GolangLines, "golang"
    case "git":
        lines, cat = GitLines, "git"
    case "docker":
        lines, cat = DockerLines, "docker"
    case "kubernetes":
        lines, cat = KubernetesLines, "kubernetes"
    case "romantic":
        lines, cat = RomanticLines, "romantic"
    case "linux":
        lines, cat = LinuxLines, "linux"
    case "networking":
        lines, cat = NetworkingLines, "networking"
    case "database":
        lines, cat = DatabaseLines, "database"
    case "cloud":
        lines, cat = CloudLines, "cloud"
    case "ai":
        lines, cat = AILines, "ai"
    case "security":
        lines, cat = SecurityLines, "security"
    case "devops":
        lines, cat = DevOpsLines, "devops"
    case "nepali":
        lines, cat = NepaliLines, "nepali"
    case "random":
        all := append(
            append(
                append(
                    append(
                        append(
                            append(
                                append(
                                    append(
                                        append(
                                            append(
                                                append(
                                                    append(
                                                        GolangLines, GitLines...),
                                                    DockerLines...),
                                                KubernetesLines...),
                                            RomanticLines...),
                                        LinuxLines...),
                                    NetworkingLines...),
                                DatabaseLines...),
                            CloudLines...),
                        AILines...),
                    SecurityLines...),
                DevOpsLines...),
            NepaliLines...,
        )
        lines, cat = all, "random"
    default:
        lines, cat = RomanticLines, "romantic"
    }

    if len(lines) == 0 {
        return cat, "No pickup lines found."
    }

    return cat, lines[rand.Intn(len(lines))]
}

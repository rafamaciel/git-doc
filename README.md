# Git Doc
Welcome to Git Doc! This extension enhances your Git workflow by providing tools to manage documentation within your repository.

## Installation
You can easily install Git Doc using curl, and we offer direct download links for various operating systems and architectures to simplify the process. Follow the instructions below to get started:

### Installation via Curl
#### Linux and macOS
To install Git Doc via curl on Linux or macOS, open your terminal and execute the appropriate command based on your architecture:

```bash
# macOS amd64
curl -Lo git-doc https://github.com/your-username/git-doc/releases/download/1.0.0/git-doc-darwin-amd64 && chmod +x git-doc && sudo mv git-doc /usr/local/bin/

# macOS arm64
curl -Lo git-doc https://github.com/your-username/git-doc/releases/download/1.0.0/git-doc-darwin-arm64 && chmod +x git-doc && sudo mv git-doc /usr/local/bin/

# Linux amd64
curl -Lo git-doc https://github.com/your-username/git-doc/releases/download/1.0.0/git-doc-linux-amd64 && chmod +x git-doc && sudo mv git-doc /usr/local/bin/

# Linux arm64
curl -Lo git-doc https://github.com/your-username/git-doc/releases/download/1.0.0/git-doc-linux-arm64 && chmod +x git-doc && sudo mv git-doc /usr/local/bin/
```

#### Windows (PowerShell)
To install Git Doc via curl on Windows using PowerShell, execute the following command:

```powershell
Invoke-WebRequest -Uri "https://github.com/your-username/git-doc/releases/latest/download/git-doc-windows-amd64.exe" -OutFile "git-doc.exe"
```

### Installation as Git Alias
Once Git Doc is installed, you can set it up as a Git alias for quick access within your repository. Run the following command in your terminal:

```bash
git config --global alias.doc '!git-doc'
```

Now you can use `git doc` to access Git Doc commands.

## Usage
Git Doc provides seamless management of documentation within your repository. Here are some examples of how you can use Git Doc:

- Initializing basic documentation structure:
  ```bash
  git doc init
  ```

- Accessing documentation through a browser:
  ```bash
  git doc serve
  ```

- Utilizing the default dcsify format for documentation.

## Contribution
We welcome contributions to Git Doc! If you'd like to contribute, please create pull requests or submit issues on our issue page.

## License
This project is licensed under the MIT License. For more details, see the LICENSE file.
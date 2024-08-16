# igen (.gitignore generator)
"igen" means "again" in Danish, which is what I say every time I am working on a webpack project. More often than not, I prefer to create a new project without it being initialized as a git repo, even if I know it will become one in the future. This may seem weird, and it is weird. I feel like creating a project with automatic git version control gives me the impression of "oh god, now I have to do this. There's no going back now. This is a real project." -- even if it's just me messing around with another language.

## My personal use cases
Say I want to work through some Advent of Code problems, or Project Euler, or anything else of the sort. A lot of the time I want to work through a solution in more than one language. Right now I am using Javascript a lot since I am learning web development, but I still love doings things in other languages. However, I do still like having access to tools like "Cargo" and Go's management system. Therefore I have two aliases in my .zshrc which create/init repos without version control. The issue with this is whenever I create a new folder for a new solution, the `.gitignore` file isn't generated since it's being initialized without git version control. 

## What it does
### NOTICE
*This functionally doesn't work yet.*

Igen checks if there's a `.gitignore` file already in the dir you're in. If there is, it will append a `.gitignore` template from whatever language you chose onto the pre-existing `.gitignore`. 


## Supported languages
* JavaScript
    - Webpack
    - Node
* Rust
* Go

## Todo
- [ ] Overwrite flag
    - [ ] If true, rename existing `.gitignore` to `.backup.gitignore`
        * create new `.gitignore` file
    - [ ] If false, append data string to existing `.gitignore`
- [x] Check CWD if `.gitignore` exists.
    - [x] If exists, append file contents of template `.gitignore` to existing file
    - [x] Else, copy template `.gitignore` from executable dir into cwd.

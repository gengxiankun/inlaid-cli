pre_prompt_command() {
    entity=$(echo $(fc -ln -1))
    [ -z "$entity" ] && return # $entity is empty or only whitespace
    $(git rev-parse --is-inside-work-tree 2> /dev/null) && local project="$(basename $(git rev-parse --show-toplevel))" || local project="Terminal"
    (~/data/code/github.com/gengxiankun/inlaid-cli/bin/inlaid-cli --path /Users/xiankun/Documents/OBSIDIAN/工作/WORKLIST.md --project "$project" --entity "$entity" 2>&1 > /dev/null &)
}

precmd_functions+=(pre_prompt_command)

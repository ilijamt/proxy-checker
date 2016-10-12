
#compdef proxy-checker
autoload -U compinit && compinit
autoload -U bashcompinit && bashcompinit

_proxy-checker_bash_autocomplete() {
    local cur prev opts base
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    opts=$( ${COMP_WORDS[0]} --completion-bash ${COMP_WORDS[@]:1:$COMP_CWORD} )
    COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
    return 0
}
complete -F _proxy-checker_bash_autocomplete proxy-checker

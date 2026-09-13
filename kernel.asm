[bits 32]
section .text
global _start

_start:
; Minimal stack setup and call into Go main
push rbp
mov rbp, rsp
sub rsp, 32                ; reserve space for locals
xor eax, eax              ; clear return value placeholder
jmp go_main_stub          ; jump to Go entry point

go_main_stub:
; Placeholder stub – actual Go code will be linked in by the build system.
; This is where we would invoke Go's runtime initialization if needed.
ret

section .data
msg db "kernel.asm stub loaded", 10, 0
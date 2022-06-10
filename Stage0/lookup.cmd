@echo 正在查找 AutoPE 文件夹
@for %%a in (C D E F G H I J K L M N O P Q R S T U V W X Y Z) do @if exist %%a:\AutoPE\ set DRIVE=%%a
@dir %DRIVE%:\AutoPE /w

md X:\AutoPE
xcopy /Y /e /i %DRIVE%:\AutoPE\stage1 X:\AutoPE\stage1

pushd X:\AutoPE
stage1\worker.exe
pause
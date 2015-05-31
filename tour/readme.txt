public repository(?)에서 package을 다운로드 받아야 할때 
	set GOPATH=C:\project-space\git-repository\go-lab\tour

git 명령어를 사용하기 위해 실행경로 설정
	set PATH=%PATH%;C:\project-space\Git\bin

code.google.com, golang.org에서 package 다운로드
	cd G:\workspace\project.lab\go2\hello
	go get -v -d code.google.com/p/go-tour/
	go get golang.org/x/tools/playground/socket
	
2015-04-10, changelog, Git 변경 사항
	C:\project-space\Git\에 설치된 git를 uninstall 하여
	http://git-scm.com/download/win 에서 다운로드 받은 파일로 설치
	
	git를 command prompt에서 사용하기 위하여 GitHub(GUI) 실행 후 'tools and options' 메뉴의
	'Open in Git Shell' 실행 (이때, 옵션에서 'Default shell'이 'Cmd'로 설정되어야, 
	go 소스	경로로 이동이 가능'
	set PATH=%PATH%;C:\project-space\Git\bin 명령어는 불필요하게 됨

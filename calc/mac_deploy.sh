#!/bin/sh

mkdir -p calc.app/Contents/Frameworks
mkdir -p calc.app/Contents/MacOS
mkdir -p calc.app/Contents/Resources

go build -o calc.app/Contents/MacOS/calc -v

otool -L calc.app/Contents/MacOS/calc

cp rc/Info.plist calc.app/Contents
cp rc/PkgInfo calc.app/Contents
cp rc/calc.icns calc.app/Contents/MacOS/Frameworks
cp -r /Library/Frameworks/Tcl.framework calc.app/Contents/Frameworks
cp -r /Library/Frameworks/Tk.framework calc.app/Contents/Frameworks

install_name_tool -change /Library/Frameworks/Tcl.framework/Versions/8.6/Tcl \
    @executable_path/../Frameworks/Tcl.framework/Versions/8.6/Tcl \
    calc.app/Contents/MacOS/calc
	
install_name_tool -change /Library/Frameworks/Tk.framework/Versions/8.6/Tk \
    @executable_path/../Frameworks/Tk.framework/Versions/8.6/Tk \
    calc.app/Contents/MacOS/calc

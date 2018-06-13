#!/bin/sh

rm -r -f qcalc.app
mkdir -p qcalc.app/Contents/Frameworks
mkdir -p qcalc.app/Contents/MacOS
mkdir -p qcalc.app/Contents/Resources

go build -o qcalc.app/Contents/MacOS/qcalc -v

otool -L qcalc.app/Contents/MacOS/qcalc

cp rc/Info.plist qcalc.app/Contents
cp rc/PkgInfo qcalc.app/Contents
cp rc/calc.icns qcalc.app/Contents/Resources
cp -p -R /Library/Frameworks/Tcl.framework qcalc.app/Contents/Frameworks
cp -p -R /Library/Frameworks/Tk.framework qcalc.app/Contents/Frameworks

install_name_tool -change /Library/Frameworks/Tcl.framework/Versions/8.6/Tcl \
    @executable_path/../Frameworks/Tcl.framework/Versions/8.6/Tcl \
    qcalc.app/Contents/MacOS/qcalc
	
install_name_tool -change /Library/Frameworks/Tk.framework/Versions/8.6/Tk \
    @executable_path/../Frameworks/Tk.framework/Versions/8.6/Tk \
    qcalc.app/Contents/MacOS/qcalc

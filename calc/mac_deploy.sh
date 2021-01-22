#!/bin/sh

rm -r -f calc.app
mkdir -p calc.app/Contents/Frameworks/Tcl.framework/Versions/8.6/Resources
mkdir -p calc.app/Contents/Frameworks/Tk.framework/Versions/8.6/Resources
mkdir -p calc.app/Contents/MacOS
mkdir -p calc.app/Contents/Resources

go build -o calc.app/Contents/MacOS/calc -v

otool -L calc.app/Contents/MacOS/calc

cp rc/Info.plist calc.app/Contents
cp rc/PkgInfo calc.app/Contents
cp rc/calc.icns calc.app/Contents/Resources

cp -p -R /Library/Frameworks/Tcl.framework calc.app/Contents/Frameworks/
cp -p -R /Library/Frameworks/Tk.framework calc.app/Contents/Frameworks/

#cp -p -R /Library/Frameworks/Tcl.framework/Versions/8.6/Resources calc.app/Contents/Frameworks/Tcl.framework/Versions/8.6/
#cp -p -R /Library/Frameworks/Tk.framework/Versions/8.6/Resources calc.app/Contents/Frameworks/Tk.framework/Versions/8.6/

#cp -p /Library/Frameworks/Tcl.framework/Versions/8.6/Tcl calc.app/Contents/Frameworks/Tcl.framework/Versions/8.6/
#cp -p /Library/Frameworks/Tcl.framework/Versions/8.6/*.sh calc.app/Contents/Frameworks/Tcl.framework/Versions/8.6/
#cp -p /Library/Frameworks/Tcl.framework/Versions/8.6/*.a calc.app/Contents/Frameworks/Tcl.framework/Versions/8.6/

#cp -p /Library/Frameworks/Tcl.framework/Tcl calc.app/Contents/Frameworks/Tcl.framework/
#cp -p /Library/Frameworks/Tcl.framework/*.sh calc.app/Contents/Frameworks/Tcl.framework/
#cp -p /Library/Frameworks/Tcl.framework/*.a calc.app/Contents/Frameworks/Tcl.framework/

#cp -p /Library/Frameworks/Tk.framework/Versions/8.6/Tk calc.app/Contents/Frameworks/Tk.framework/Versions/8.6/
#cp -p /Library/Frameworks/Tk.framework/Versions/8.6/*.sh calc.app/Contents/Frameworks/Tk.framework/Versions/8.6/
#cp -p /Library/Frameworks/Tk.framework/Versions/8.6/*.a calc.app/Contents/Frameworks/Tk.framework/Versions/8.6/

#cp -p /Library/Frameworks/Tk.framework/Tk calc.app/Contents/Frameworks/Tk.framework/
#cp -p /Library/Frameworks/Tk.framework/*.sh calc.app/Contents/Frameworks/Tk.framework/
#cp -p /Library/Frameworks/Tk.framework/*.a calc.app/Contents/Frameworks/Tk.framework/

#cp -p /Library/Frameworks/Tcl.framework/Versions/8.6/* calc.app/Contents/Frameworks/Tcl.framework/Versions/8.6/
#cp -p /Library/Frameworks/Tk.framework/Versions/8.6/* calc.app/Contents/Frameworks/Tk.framework/Versions/8.6/
#cp -p /Library/Frameworks/Tcl.framework/Tcl calc.app/Contents/Frameworks/Tcl.framework/Tcl
#cp -p /Library/Frameworks/Tk.framework/Versions/8.6/* calc.app/Contents/Frameworks/Versions/8.6/

install_name_tool -change /Library/Frameworks/Tcl.framework/Versions/8.6/Tcl \
    @executable_path/../Frameworks/Tcl.framework/Versions/8.6/Tcl \
    calc.app/Contents/MacOS/calc
	
install_name_tool -change /Library/Frameworks/Tk.framework/Versions/8.6/Tk \
    @executable_path/../Frameworks/Tk.framework/Versions/8.6/Tk \
    calc.app/Contents/MacOS/calc

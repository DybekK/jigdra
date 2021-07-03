import 'package:flutter/material.dart';
import 'package:flutter_spinkit/flutter_spinkit.dart';
import 'package:jigdra/jigdra_packages_provider.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
        child: Column(mainAxisAlignment: MainAxisAlignment.center, children: [
          AutoSizeText('Hello'),
          SpinKitFadingFour(
            duration: Duration(seconds: 1),
            color: Colors.black,
          ),
        ]),
      ),
    );
  }
}

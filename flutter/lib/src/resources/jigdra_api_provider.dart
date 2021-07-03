import 'dart:convert';

import 'package:http/http.dart' as http;
import '../models/models.dart';

class LoginRequestFailure implements Exception {}

class JigdraApiClientProvider {
  JigdraApiClientProvider({http.Client? httpClient})
      : _httpClient = httpClient ?? http.Client();

  static const _baseUrl = "www.example.com.pl";
  final http.Client _httpClient;

  //TODO implement request methods
}

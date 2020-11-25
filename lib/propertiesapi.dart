import 'dart:convert';
import 'dart:html';

import 'package:mango_ui/keys.dart';
import 'package:mango_ui/requester.dart';

import 'bodies/property.dart';

Future<HttpRequest> createProperty(Property obj) async {
  var apiroute = getEndpoint("house");
  var url = "${apiroute}/info";

  return invokeService("POST", url, jsonEncode(obj.toJson()));
}

Future<HttpRequest> updateProperty(Key key, Property obj) async {
  var route = getEndpoint("house");
  var url = "${route}/info/${key.toJson()}";

  final data = jsonEncode(obj.toJson());

  return invokeService("PUT", url, data);
}

Future<HttpRequest> deleteProperty(Key key) async {
  var route = getEndpoint("house");
  var url = "${route}/info/${key.toJson()}";

  return invokeService("DELETE", url, "");
}

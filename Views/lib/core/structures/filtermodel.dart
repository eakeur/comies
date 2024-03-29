abstract class Filter {
  int? skip;

  Map<String, dynamic> toMap();

  Filter.fromMap(Map<String, dynamic> map);

  String toJson();

  Filter.fromJson(String source);

}

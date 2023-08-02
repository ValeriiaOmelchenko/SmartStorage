int data;
void setup() {
  // put your setup code here, to run once:
 Serial.begin(9600);
  Serial3.begin(9600);

}

void loop() {
  // put your main code here, to run repeatedly:
if (Serial3.available()) {
    data = Serial3.parseInt();
    Serial.println(data);
    delay(10000);
}
}
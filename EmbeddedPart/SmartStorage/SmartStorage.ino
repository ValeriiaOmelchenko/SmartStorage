int data;
byte currentLED;
byte minLED=2;
byte maxLED=5;
void setup() {
  // put your setup code here, to run once:
  Serial.begin(9600);
  Serial3.begin(9600);
  for (byte i=minLED;i<=maxLED;i++){
      pinMode(i, OUTPUT);
  }
}

void loop() {
  // put your main code here, to run repeatedly:
  if (Serial3.available()) {
    data = Serial3.parseInt();
    if((data>=minLED)&&(data<=maxLED)||(data==0)){
    if ((currentLED != data)&&(data!=0)) {
digitalWrite(currentLED, LOW);
      currentLED = data;
    }
    Serial.println(data);
    digitalWrite(data, HIGH);
    delay(1000);
    }
    else{
      Serial3.print(99);
    }
  }
}
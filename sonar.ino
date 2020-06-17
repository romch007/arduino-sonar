#include <Servo.h>

#define BAUD_RATE 9600
#define SERVO_PIN 7
#define TRIG_PIN 6
#define ECHO_PIN 5

#define ANGLE_STEP 5

Servo servo;

/* Servo :
 * Left : 170° 
 * Right: 0°
 */

bool in_scan;
int current_angle;
int current_distance;

int compute_distance() {
  int distance;
  float duration;

  digitalWrite(TRIG_PIN, LOW);
  delayMicroseconds(2);

  digitalWrite(TRIG_PIN, HIGH);
  delayMicroseconds(10);
  digitalWrite(TRIG_PIN, LOW);

  duration = pulseIn(ECHO_PIN, HIGH);
  distance = duration * 0.034 / 2;

  return distance;
}

void send_infos(int angle, int distance) {
  Serial.print(angle, DEC);
  Serial.print(",");
  Serial.print(distance, DEC);
  Serial.print("\n");
}

void setup() {
  Serial.begin(BAUD_RATE);
  pinMode(TRIG_PIN, OUTPUT);
  pinMode(ECHO_PIN, INPUT);  
  in_scan = true;
  current_angle = 0;
  servo.attach(SERVO_PIN);
}

void loop() {
  String command = Serial.readString();

  if (command == "stop\n") {
    in_scan = false;
  }

  if (current_angle >= 170) {
    in_scan = false;
    Serial.print("end\n");
    servo.write(85);
    delay(10000);
  }

  if (in_scan) {
    servo.write(current_angle);
    current_distance = compute_distance();
    send_infos(current_angle, current_distance);
    current_angle += ANGLE_STEP;
    delay(100);
  }
}

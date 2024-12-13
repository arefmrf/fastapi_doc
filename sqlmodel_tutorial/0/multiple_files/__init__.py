from user.models import User
from django.core.files import File
from host.models import HostType, Place, Host, HostImage


HostType.save = lambda self, *args, **kwargs: super(HostType, self).save(*args, **kwargs)
host_type = HostType.objects.create(category=1, title='room', translate={"ir": "خانه", "en": "room"})
host_type.icon.save('batman1.jpeg', File(open('batman1.jpeg', 'rb')))
user = User.objects.first()

for i in range(1, 800):
    place = Place.objects.create(country="country", province="province", city="city", address="addressaddressaddress", lat=0.1, lng=0.2)
    host = Host.objects.create(place=place, host_type=host_type, user=user.id, name="name", status=Host.Status.ACTIVE)
    image = HostImage.objects.create(host=host)
    image.image.save('batman1.jpeg', File(open('batman1.jpeg', 'rb')))

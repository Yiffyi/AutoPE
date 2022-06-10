using System.Collections.Generic;
using System.Linq;
using System.Management;

namespace AutoPE.WMI
{
    public class WMIObject
    {
        public ManagementObject RawObject;

        static public IEnumerable<ManagementObject> FromWQL(string query)
        {
            ManagementObjectSearcher searcher = new ManagementObjectSearcher(query);
            return searcher.Get()
                .Cast<ManagementObject>();
        }

        static public IEnumerable<ManagementObject> GetInstances(string className)
        {
            ManagementClass mc = new ManagementClass(className);
            return mc.GetInstances().Cast<ManagementObject>();
        }

        public WMIObject(ManagementObject rawObject)
        {
            RawObject = rawObject;
        }


        public object this[string i]
        {
            get => RawObject?[i];
            set => RawObject.SetPropertyValue(i, value);
        }

        public ManagementPath CommitChanges() => RawObject.Put();

        public T CallMethod<T>(string method, object input = null)
        {
            var inParams = RawObject.GetMethodParameters(method);

            var props = input.GetType().GetProperties();
            foreach (var prop in props)
            {
                inParams[prop.Name] = prop.GetValue(input);
            }

            var outParams = RawObject.InvokeMethod(method, inParams, null);
            return (T)outParams.Properties["ReturnValue"].Value;
        }
    }
}
